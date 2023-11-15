# STIR/SHAKEN Certificate Repository Compliance

## Encore Communications

Name: `https://cdn-cr.cgah.tnsi.com/certs/a1b25708d639f93a043a23170376e6bc24aea56c`\
Tested At: 15 Nov 23 16:09 UTC\
Time: 71ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 15 Nov 23 17:17 UTC