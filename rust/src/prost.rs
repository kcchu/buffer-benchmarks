use std::io::Cursor;

use prost::Message;

use crate::constants;

mod generated {
    include!(concat!(env!("OUT_DIR"), "/generated.rs"));
}

pub fn encode() -> Vec<u8> {
    let md = generated::MessageData {
        body: Some(generated::message_data::Body::CastAddBody(
            generated::CastAddBody {
                embeds: Default::default(),
                mentions: Default::default(),
                text: String::from(constants::SAMPLE_TEXT),
                parent: Some(generated::cast_add_body::Parent::CastId(
                    generated::CastId {
                        fid: Vec::from(constants::SAMPLE_FID),
                        ts_hash: Vec::from(constants::SAMPLE_TS_HASH),
                    },
                )),
            },
        )),
        r#type: generated::MessageType::CastAdd.into(),
        timestamp: constants::SAMPLE_TIMESTAMP,
        fid: Vec::from(constants::SAMPLE_FID),
        network: generated::FarcasterNetwork::Devnet.into(),
    };
    let mut md_buf = Vec::new();
    md_buf.reserve(md.encoded_len());
    md.encode(&mut md_buf).unwrap();

    let m = generated::Message {
        data: md_buf,
        hash: Vec::from(constants::SAMPLE_HASH),
        hash_scheme: generated::HashScheme::Blake3.into(),
        signature: Vec::from(constants::SAMPLE_SIGNATURE),
        signature_scheme: generated::SignatureScheme::Ed25519.into(),
        signer: Vec::from(constants::SAMPLE_SIGNER),
    };
    let mut m_buf = Vec::new();
    m_buf.reserve(m.encoded_len());
    m.encode(&mut m_buf).unwrap();
    m_buf.into()
}

pub fn decode(buf: &[u8]) {
    let m = generated::Message::decode(Cursor::new(buf)).unwrap();
    let md = generated::MessageData::decode(Cursor::new(m.data)).unwrap();
    let generated::message_data::Body::CastAddBody(body) = md.body.unwrap();
    if body.text != constants::SAMPLE_TEXT {
        panic!("Unexpected decoded text")
    }
}

#[cfg(test)]
mod tests {
    use test::bench::Bencher;

    #[bench]
    fn bench_encode(b: &mut Bencher) {
        b.iter(|| _ = super::encode());
    }

    #[bench]
    fn bench_decode(b: &mut Bencher) {
        let buf = super::encode();
        println!("Wire format size (bytes) = {}", buf.len());
        b.iter(|| _ = super::decode(&buf));
    }
}