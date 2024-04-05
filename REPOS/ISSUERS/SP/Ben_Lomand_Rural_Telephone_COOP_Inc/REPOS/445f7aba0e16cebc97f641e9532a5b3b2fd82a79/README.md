# STIR/SHAKEN Certificate Repository Compliance

## Ben Lomand Rural Telephone COOP Inc

Name: `https://cdn-cr.cgah.tnsi.com/certs/fe74899c830827a37cd701d031e0a7714e0604ab`\
Tested At: 05 Apr 24 18:39 UTC\
Time: 30ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 05 Apr 24 19:04 UTC