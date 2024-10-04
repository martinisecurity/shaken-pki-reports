# STIR/SHAKEN Certificate Repository Compliance

## Business Telecommunications Services Inc

Name: `https://stirshaken-m.bts.io/api/v1/certificates/bts_b3f8527d0c31ee0d9d9e19da01e8cf68.cer`\
Tested At: 04 Oct 24 16:26 UTC\
Time: 327ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC