# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

Name: e_atis_signature_algorithm_ca\
Source: ATIS1000080\
Citation: ATIS-1000080.v003 / 6.4.1 SHAKEN Certificate Requirements\
Effective Date: 04 Sep 20 00:00 UTC\
Description: STI certificates shall contain a Signature Algorithm field with the value 'ecdsa-with-SHA256'

### Leaf Certificates

No error, warning, or notice level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | NetNumber SHAKEN Root Intermediate CA 1 | [view](../../CERTS/e449803766edf02ab50b034dd7e89e54efd332cce87688a032f89b340d039878/README.md) | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.10045.4.3.3 |
| error | NetNumber SHAKEN Root CA 1 | [view](../../CERTS/2dd1386ca717f31d550b35b9bce9daa9b02483bcdb98bdfcfca07202276136d7/README.md) | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.10045.4.3.4 |


Generated: 24 Nov 23 11:17 UTC