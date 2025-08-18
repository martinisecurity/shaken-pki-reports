# STIR/SHAKEN Certificate Repository Compliance

## Crosstel Tandem Inc SEPB

Name: `https://cdn-cr.cgah.tnsi.com/certs/0561e55091b73ef2ead98e0b0ba52f1762fffa9a`\
Tested At: 18 Aug 25 20:05 UTC\
Time: 21ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC