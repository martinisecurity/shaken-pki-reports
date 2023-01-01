# STIR/SHAKEN Certificate Repository Compliance

## Bandwidth.com CLEC LLC

Name: `https://bw-shaken-cert-pub.s3.amazonaws.com/bandwidth-shaken-cert_20230716.pem`\
Tested At: 01 Jan 23 23:23 UTC\
Time: 50ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01 Jan 23 23:34 UTC