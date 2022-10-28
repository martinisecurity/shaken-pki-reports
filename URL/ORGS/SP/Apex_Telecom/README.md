# STIR/SHAKEN Certificate Repository Compliance

## Apex Telecom

| Code | Source | Instances |
|------|--------|-----------|
| [w_atis_content_type](ISSUES/w_atis_content_type/README.md) | ATIS-1000080 | 1 |
| [e_atis_cache_header](ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 | 1 |

### https://s3.amazonaws.com/certificates.peeringhub.io/288K/288K.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |


Generated: 28/10/2022 at 18:22:16