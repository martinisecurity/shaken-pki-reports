# STIR/SHAKEN Certificate Repository Compliance

## Ribbon Communications

Name: `https://prod001-prod011-cr.rbbnidhub.com/R3kZfzj7gz/STI-202306-2130-3e01a84e6a320c2e0d9baf5cbeb11f38`\
Tested At: 22 Aug 24 15:36 UTC\
Time: 351ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_http_status_200](../../ISSUES/e_http_status_200/README.md) | error | HTTP | HTTP response shall have StatusCode 200, but it is 403 Forbidden |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 22 Aug 24 15:44 UTC