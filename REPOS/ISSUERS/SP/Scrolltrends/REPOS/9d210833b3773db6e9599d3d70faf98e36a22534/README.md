# STIR/SHAKEN Certificate Repository Compliance

## Scrolltrends

Name: `http://certs.icallr.com/551L/scrolltrends.pem`\
Tested At: 18 Aug 25 20:04 UTC\
Time: 487ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_atis_redirect](../../ISSUES/e_atis_redirect/README.md) | error | ATIS-1000074 | The STI-VS shall not follow HTTP redirections |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC