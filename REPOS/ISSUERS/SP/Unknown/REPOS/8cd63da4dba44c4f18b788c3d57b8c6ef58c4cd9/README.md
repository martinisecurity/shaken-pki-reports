# STIR/SHAKEN Certificate Repository Compliance

## Unknown

Name: `https://api.alianza.com/v2/stir-shaken/certs/b45a4083-1554-4412-b5fc-bbd2c027091e/key.crt`\
Tested At: 31 Oct 22 19:19 UTC\
Time: 511ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 31/10/2022 at 19:21:49