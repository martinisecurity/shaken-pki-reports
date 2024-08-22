# STIR/SHAKEN Certificate Repository Compliance

## Ribbon Communications

Name: `https://prod001-prod011-cr.rbbnidhub.com/o_W93Cgngz/STI-202307-3390-7148553e5c7080fef0ac81ea7235258e`\
Tested At: 22 Aug 24 15:59 UTC\
Time: 56ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_http_status_200](../../ISSUES/e_http_status_200/README.md) | error | HTTP | HTTP response shall have StatusCode 200, but it is 403 Forbidden |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 22 Aug 24 16:06 UTC