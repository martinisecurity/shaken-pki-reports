# STIR/SHAKEN Certificate Repository Compliance

## Unknown

Name: `http://5.161.181.254/ec256-public.pem`\
Tested At: 18 Aug 25 20:04 UTC\
Time: 27ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_http_status_200](../../ISSUES/e_http_status_200/README.md) | error | HTTP | HTTP response shall have StatusCode 200, but it is 404 Not Found |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://5.161.181.254/ec256-public.pem": x509: cannot validate certificate for 5.161.181.254 because it doesn't contain any IP SANs |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 18 Aug 25 21:13 UTC