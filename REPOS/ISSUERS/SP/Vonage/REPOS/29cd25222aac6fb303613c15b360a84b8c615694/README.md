# STIR/SHAKEN Certificate Repository Compliance

## Vonage

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/56701755b64fd6472ffc065ab0dfe768.pem`\
Tested At: 21 Nov 23 17:10 UTC\
Time: 612ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Nov 23 17:16 UTC