# STIR/SHAKEN Certificate Repository Compliance

## Vonage

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/56701755b64fd6472ffc065ab0dfe768.pem`\
Tested At: 12 Feb 24 16:57 UTC\
Time: 85ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 17:02 UTC