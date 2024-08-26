# STIR/SHAKEN Certificate Repository Compliance

## Cspire

Name: `https://cdn-cr.cgah.tnsi.com/certs/2487f31aa2e09ad03fa91fa776793115d03392e9`\
Tested At: 26 Aug 24 17:42 UTC\
Time: 149ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:03 UTC