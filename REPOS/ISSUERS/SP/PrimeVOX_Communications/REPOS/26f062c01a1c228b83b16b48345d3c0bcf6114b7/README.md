# STIR/SHAKEN Certificate Repository Compliance

## PrimeVOX Communications

Name: `https://pvx1.s3.us-east-2.amazonaws.com/stirshaken/8448cc7eb8424d6ad5d2e6d71bcf6629.cer`\
Tested At: 01 Nov 22 16:29 UTC\
Time: 380ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01/11/2022 at 16:30:07