# STIR/SHAKEN Certificate Repository Compliance

## Unknown

Name: `https://www.hadlotechnologies.com/ss_certs/hadlo_stirshaken.public.crt`\
Tested At: 12 Feb 24 19:22 UTC\
Time: 188ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header doesn't have the 'private' tag |
| [e_http_status_200](../../ISSUES/e_http_status_200/README.md) | error | HTTP | HTTP response shall have StatusCode 200, but it is 404 Not Found |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://www.hadlotechnologies.com/ss_certs/hadlo_stirshaken.public.crt": x509: certificate has expired or is not yet valid: current time 2024-02-12T19:22:37Z is after 2024-02-11T23:59:59Z |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 12 Feb 24 19:26 UTC