# STIR/SHAKEN Certificate Repository Compliance

## Twin Valley

Name: `https://cdn-cr.cgah.tnsi.com/certs/f292d2a556a66a810fed9e79097db082bd9a66a2`\
Tested At: 18 Aug 25 20:05 UTC\
Time: 20ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC