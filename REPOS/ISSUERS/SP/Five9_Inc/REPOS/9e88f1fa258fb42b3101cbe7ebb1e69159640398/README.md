# STIR/SHAKEN Certificate Repository Compliance

## Five9 Inc

Name: `https://sticr-cstga-uat.ccid.neustar/api/v1/certificate/d1dbe1a693cb7464ff1d1edb3fee78e2.pem`\
Tested At: 04 Oct 24 16:25 UTC\
Time: 348ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC