# STIR/SHAKEN Certificate Repository Compliance

## 9171-5573 Quebec Inc

Name: `https://shaken-9171-5573-quebec-inc.s3.us-east-2.amazonaws.com/pubKeys/4722b9c5.crt`\
Tested At: 18 Aug 25 21:05 UTC\
Time: 181ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC