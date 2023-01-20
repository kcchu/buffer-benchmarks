mod message_generated;

use crate::constants;

use self::message_generated::farcaster as generated;

pub fn encode() -> Vec<u8> {
    let mut mdb = flatbuffers::FlatBufferBuilder::with_capacity(1024);

    let text = mdb.create_string(constants::SAMPLE_TEXT);
    let fid = mdb.create_vector(&constants::SAMPLE_FID);
    let parent_fid = mdb.create_vector(&constants::SAMPLE_FID);
    let parent_ts_hash = mdb.create_vector(&constants::SAMPLE_TS_HASH);

    // CastId
    let parent = generated::CastId::create(
        &mut mdb,
        &generated::CastIdArgs {
            fid: Some(parent_fid),
            ts_hash: Some(parent_ts_hash),
        },
    );

    // CastAddBody
    let cast_add_body = generated::CastAddBody::create(
        &mut mdb,
        &generated::CastAddBodyArgs {
            embeds: None,
            mentions: None,
            parent_type: generated::TargetId::CastId,
            parent: Some(parent.as_union_value()),
            text: Some(text),
        },
    );

    // MessageData
    let message_data = generated::MessageData::create(
        &mut mdb,
        &generated::MessageDataArgs {
            body_type: generated::MessageBody::CastAddBody,
            body: Some(cast_add_body.as_union_value()),
            type_: Some(generated::MessageType::CastAdd),
            timestamp: constants::SAMPLE_TIMESTAMP,
            fid: Some(fid),
            network: generated::FarcasterNetwork::Devnet,
        },
    );

    mdb.finish(message_data, None);

    let mut mb = flatbuffers::FlatBufferBuilder::with_capacity(1024);
    let data = mb.create_vector(mdb.finished_data());
    let hash = mb.create_vector(&constants::SAMPLE_HASH);
    let signature = mb.create_vector(&constants::SAMPLE_SIGNATURE);
    let signer = mb.create_vector(&constants::SAMPLE_SIGNER);

    // Message
    let message = generated::Message::create(
        &mut mb,
        &generated::MessageArgs {
            data: Some(data),
            hash: Some(hash),
            hash_scheme: generated::HashScheme::Blake3,
            signature: Some(signature),
            signature_scheme: generated::SignatureScheme::Ed25519,
            signer: Some(signer),
        },
    );

    mb.finish(message, None);
    Vec::from(mb.finished_data())
}

pub fn decode(buf: &[u8]) {
    let m = generated::root_as_message(buf).unwrap();
    let data = m.data();
    let md = flatbuffers::root::<generated::MessageData>(data.bytes()).unwrap();
    let body = md.body_as_cast_add_body().unwrap();
    if body.text() != constants::SAMPLE_TEXT {
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
