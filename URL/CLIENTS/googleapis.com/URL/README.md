# STIR/SHAKEN Certificate Repository Compliance

## googleapis.com

| Code | Source | Instances |
|------|--------|-----------|
| [e_atis_cache_header](../ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 | 1 |
| [w_atis_content_type](../ISSUES/w_atis_content_type/README.md) | ATIS-1000080 | 1 |

### https://storage.googleapis.com/stirshaken/ShakeNBakeCert.pem

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The expiration time is less than 24 hours |


Generated: 27/10/2022 at 21:42:12