# STIR/SHAKEN Certificate Repository Compliance

## Pulsar360 Corp

Name: `https://ecms.securetransit.net/certs/06517b2e2010ec1a2b429664047353b6.cer`\
Tested At: 22 Aug 24 15:56 UTC\
Time: 460ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 16:06 UTC