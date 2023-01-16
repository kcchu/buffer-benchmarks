using Go = import "/go.capnp";
@0xe6989e0b6c84e686;

$Go.package("generated");
$Go.import("github.com/kcchu/buffer-benchmarks/capnproto/generated");

enum MessageType {
  defaultMessageType @0;
  castAdd @1;
  castRemove @2;
  reactionAdd @3;
  reactionRemove @4;
  ampAdd @5;
  ampRemove @6;
  verificationAddEthAddress @7;
  verificationRemove @8;
  signerAdd @9;
  signerRemove @10;
  userDataAdd @11;
}

enum SignatureScheme {
  defaultSignatureScheme @0;
  ed25519 @1;
  eip712 @2;
}

enum HashScheme {
  defaultHashScheme @0;
  blake3 @1;
}

enum FarcasterNetwork {
  defaultFarcasterNetwork @0;
  mainnet @1;
  testnet @2;
  devnet @3;
}

struct CastId {
  fid @0 :Data;
  tsHash @1 :Data;
}

struct UserId {
  fid @0 :Data;
}

struct CastAddBody {
  embeds @0 :List(Text);
  mentions @1 :List(UserId);
  parent :union {
    void @2 :Void;
    castId @3 :CastId;
  }
  text @4 :Text;
}

struct MessageData {
  body :union {
    void @0 :Void;
    castAddBody @1 :CastAddBody;
  }
  type @2 :MessageType;
  timestamp @3 :UInt32;
  fid @4 :Data;
  network @5 :FarcasterNetwork;
}

struct Message {
  data @0 :Data;
  hash @1 :Data;
  hashScheme @2 :HashScheme;
  signature @3 :Data;
  signatureScheme @4 :SignatureScheme;
  signer @5 :Data;
}