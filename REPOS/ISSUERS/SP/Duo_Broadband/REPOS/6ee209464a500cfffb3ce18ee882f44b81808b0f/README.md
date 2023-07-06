# STIR/SHAKEN Certificate Repository Compliance

## Duo Broadband

Name: `https://cdn-cr.cgah.tnsi.com/certs/54a54551e139f777dcad69a463d282d8b6b9dade`\
Tested At: 06 Jul 23 13:53 UTC\
Time: 23ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 06 Jul 23 14:08 UTC