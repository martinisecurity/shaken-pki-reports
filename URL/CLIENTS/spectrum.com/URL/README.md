# STIR/SHAKEN Certificate Repository Compliance

## spectrum.com

| Code | Source | Instances |
|------|--------|-----------|
| [w_atis_content_type](../ISSUES/w_atis_content_type/README.md) | ATIS-1000080 | 2 |
| [e_atis_cache_header](../ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 | 2 |

### https://shaken.spectrum.com/4d65efdb8a1ca366e9576c8fda747fa4.pem

| Code | Status | Source | Details |
|------|--------|--------|---------|
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

### https://shaken.spectrum.com/cf1b3d3d-7f2b-42fd-a161-ebe61cd6565a.pem

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |


Generated: 27/10/2022 at 21:25:02