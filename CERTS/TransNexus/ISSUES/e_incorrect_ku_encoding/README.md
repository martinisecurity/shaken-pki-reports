# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

Name: e_incorrect_ku_encoding\
Source: RFC5280\
Citation: Where ITU-T Rec. X.680 | ISO/IEC 8824-1, 21.7, applies, the bitstring shall have all trailing 0 bits removed before it is encoded.\
Effective Date: 01 Jan 00 00:00 UTC\
Description: RFC 5280 Section 4.2.1.3 describes the value of a KeyUsage to be a DER encoded BitString, which itself defines that all trailing 0 bits be counted as being "unused".

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | SHAKEN 031K | [view](../../CERTS/a6b43e7cf487f1fd10f20b7197a556d8aecab5c71e91ab87a634894ca4fb00f6/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 066K | [view](../../CERTS/51c8cf0bbfb640f7ef4e7e47bedd1a9a1d4f757fe142ad05fa84422544794ccd/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 073H | [view](../../CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 115K | [view](../../CERTS/631861bcfeb6aeaf83a56e5c1c376c8981927a52fc8da57964b1776b69e2bf6e/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 166K | [view](../../CERTS/d4d723c554cc42a2cd242010ffd2499aab6436955e8031d4f3109db3b66a3b64/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 193E | [view](../../CERTS/2a2a297de4aa7620b8f38d5266dd4f76fbd6b5c79a87411b589c2226518ce475/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 345J | [view](../../CERTS/27c018db560aefbff83f32326db2cbf3f1260051386569d7775cf98fae145bae/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 345J | [view](../../CERTS/109a4ac394ec4e6eb6ba42d2ad6bc799f2548629a8bbfb9420b9ec1ca5df6ff9/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 505J | [view](../../CERTS/b186c959554e652b69824a2e45ac08ec135105b0e01c9d2eadeb3cf46c130670/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 505J | [view](../../CERTS/1c0003ac10eeaa04229e507c15f71ee018c3902c3e8c20fbe42e533b8682ba8c/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 578J | [view](../../CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 597J | [view](../../CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 622J | [view](../../CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 6628 | [view](../../CERTS/37b9a67bb8e3272048330f269b5dbc285ac0278df4670ab7e291b92b6548fa2d/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | CCI SHAKEN 663J | [view](../../CERTS/9cfd0b9974acdc9b5fce3888cf2f6cf99173e7955adeebd37a1e60dfc444a0b3/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 706J | [view](../../CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 722J | [view](../../CERTS/50eb0c6a670f4122f8fbdb75582aa257fe3979f441fb396ff738372627104f9c/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 736J | [view](../../CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 755J | [view](../../CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 815G | [view](../../CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 8526 | [view](../../CERTS/b68dadcfa6267b8e9dd012e42154b55cab6d3694f4543e46af44bab1e4ba971e/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | MobileSphere SHAKEN 873J | [view](../../CERTS/8f74ddd5e6964042b119d5b3f17d24694ec78c8688e0ffc96d13fdb0e8df26d4/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 952J | [view](../../CERTS/bbccd67a84881c78b37ecfb961dd7029498490975d91e15f872d5a684a225626/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | SHAKEN 857J | [view](../../CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| error | Fusion Connect SHAKEN 2720 | [view](../../CERTS/03a010b294807e90cc1309d7466fd2ddc591158a69be6950ff4d0eab32de0799/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | TransNexus, Inc. SHAKEN Issuing CA3 | [view](../../CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |
| error | TransNexus, Inc. SHAKEN Issuing CA1 | [view](../../CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |
| error | TransNexus, Inc. SHAKEN Root CA1 | [view](../../CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |


Generated: 04 Nov 22 01:11 UTC