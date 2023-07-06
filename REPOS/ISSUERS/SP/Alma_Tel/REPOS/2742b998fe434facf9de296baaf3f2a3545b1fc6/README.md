# STIR/SHAKEN Certificate Repository Compliance

## Alma Tel

Name: `https://cdn-cr.cgah.tnsi.com/certs/50fce21c2814c12146d9f3e9420b586cd6c12566`\
Tested At: 06 Jul 23 13:53 UTC\
Time: 12ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 06 Jul 23 14:08 UTC