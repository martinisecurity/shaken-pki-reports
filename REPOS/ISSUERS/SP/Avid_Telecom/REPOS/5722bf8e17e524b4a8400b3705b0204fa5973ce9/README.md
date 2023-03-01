# STIR/SHAKEN Certificate Repository Compliance

## Avid Telecom

Name: `https://ssc.getsipnav.com/certs/d29187037eabb3b60e24baba5c0b3dc424e004b9`\
Tested At: 01 Mar 23 18:18 UTC\
Time: 85ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01 Mar 23 18:22 UTC