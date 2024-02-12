# STIR/SHAKEN Certificate Repository Compliance

## Cloudli Communications

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/7bdca6d62828d732c42a364bd6e6662a.pem`\
Tested At: 12 Feb 24 16:57 UTC\
Time: 278ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 17:02 UTC