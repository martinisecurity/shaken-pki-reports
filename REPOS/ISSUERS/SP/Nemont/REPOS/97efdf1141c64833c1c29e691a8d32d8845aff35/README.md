# STIR/SHAKEN Certificate Repository Compliance

## Nemont

Name: `https://cdn-cr.cgah.tnsi.com/certs/da7847867ee00785d849b849e374c81ecd3f2545`\
Tested At: 21 Nov 23 17:37 UTC\
Time: 22ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Nov 23 17:53 UTC