# STIR/SHAKEN Certificate Repository Compliance

## TekSavvy SHAKEN 548H

Name: `https://netnumber-sti-cr.s3.amazonaws.com/certs/bbd520f5-874f-4a8f-9408-a09e2e0e76ea`\
Tested At: 05 Apr 24 18:53 UTC\
Time: 199ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 05 Apr 24 19:04 UTC