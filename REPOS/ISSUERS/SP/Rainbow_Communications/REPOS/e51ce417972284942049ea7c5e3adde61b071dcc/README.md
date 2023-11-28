# STIR/SHAKEN Certificate Repository Compliance

## Rainbow Communications

Name: `https://cdn-cr.cgah.tnsi.com/certs/861edae589b34de5d112e3448589fd6c3d346e94`\
Tested At: 28 Nov 23 19:52 UTC\
Time: 30ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Nov 23 20:21 UTC