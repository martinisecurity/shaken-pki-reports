# STIR/SHAKEN Certificate Repository Compliance

## accessline.com

| Code | Source | Instances |
|------|--------|-----------|
| [w_atis_content_type](../ISSUES/w_atis_content_type/README.md) | ATIS-1000080 | 1 |
| [e_atis_cache_header](../ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 | 1 |

### https://sscerts.accessline.com/sscerts/Cert-US-04-2022.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |


Generated: 27/10/2022 at 22:42:50