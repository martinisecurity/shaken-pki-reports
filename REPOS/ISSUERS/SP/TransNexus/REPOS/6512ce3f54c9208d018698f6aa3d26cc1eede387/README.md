# STIR/SHAKEN Certificate Repository Compliance

## TransNexus

Name: `https://shaken.spectrum.com/cf1b3d3d-7f2b-42fd-a161-ebe61cd6565a.pem`\
Tested At: 30 Nov 22 17:23 UTC\
Time: 40ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 30 Nov 22 17:24 UTC