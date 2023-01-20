mod codegen {
    include!(concat!(env!("OUT_DIR"), "/protobuf/mod.rs"));
}

use protobuf::Message;

use self::codegen::message as generated;
use crate::constants;

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
                        special_fields: Default::default(),
                    },
                )),
                special_fields: Default::default(),
            },
        )),
        type_: generated::MessageType::CastAdd.into(),
        timestamp: constants::SAMPLE_TIMESTAMP,
        fid: Vec::from(constants::SAMPLE_FID),
        network: generated::FarcasterNetwork::Devnet.into(),
        special_fields: Default::default(),
    };
    let md_buf = md.write_to_bytes().unwrap();

    let m = generated::Message {
        data: md_buf,
        hash: Vec::from(constants::SAMPLE_HASH),
        hash_scheme: generated::HashScheme::Blake3.into(),
        signature: Vec::from(constants::SAMPLE_SIGNATURE),
        signature_scheme: generated::SignatureScheme::Ed25519.into(),
        signer: Vec::from(constants::SAMPLE_SIGNER),
        special_fields: Default::default(),
    };
    m.write_to_bytes().unwrap()
}

pub fn decode(buf: &[u8]) {
    let m = generated::Message::parse_from_bytes(buf).unwrap();
    let md = generated::MessageData::parse_from_bytes(m.data.as_ref()).unwrap();
    let generated::message_data::Body::CastAddBody(body) = md.body.unwrap();
    if body.text != constants::SAMPLE_TEXT {
        panic!("Unexpected decoded text")
    }
}
