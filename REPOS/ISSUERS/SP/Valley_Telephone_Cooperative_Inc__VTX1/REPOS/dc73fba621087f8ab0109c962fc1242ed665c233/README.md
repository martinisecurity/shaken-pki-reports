# STIR/SHAKEN Certificate Repository Compliance

## Valley Telephone Cooperative Inc / VTX1

Name: `https://cdn-cr.cgah.tnsi.com/certs/4cd5b91a833e4ccd9c7f4dabacdd26e85cddf71d`\
Tested At: 26 Aug 24 17:42 UTC\
Time: 33ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:03 UTC