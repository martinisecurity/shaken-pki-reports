# STIR/SHAKEN Certificate Repository Compliance

## Unknown

Name: `http://stirshaken.ellianz.com/`\
Tested At: 04 Oct 24 15:30 UTC\
Time: 919ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Pragma header contains 'no-cache' |
| [e_atis_redirect](../../ISSUES/e_atis_redirect/README.md) | error | ATIS-1000074 | The STI-VS shall not follow HTTP redirections |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 04 Oct 24 16:29 UTC