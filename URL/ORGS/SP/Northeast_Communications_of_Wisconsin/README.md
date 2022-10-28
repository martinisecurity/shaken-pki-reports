# STIR/SHAKEN Certificate Repository Compliance

## Northeast Communications of Wisconsin

| Code | Source | Instances |
|------|--------|-----------|
| [e_atis_cache_header](ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 | 1 |
| [w_atis_content_type](ISSUES/w_atis_content_type/README.md) | ATIS-1000080 | 1 |

### https://cdn-cr.cgah.tnsi.com/certs/1a7ab760b70a56725786e626f157ba4512f909ec

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |


Generated: 28/10/2022 at 19:21:41