# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN OneStream Networks, LLC 630J

Tested At: 24 Nov 23 11:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 356 day(s)\
Subject: CN=SHAKEN OneStream Networks\\, LLC 630J, OU=OneStream Networks, O=OneStream Networks\\, LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/OneStream_Networks_LLC_630J

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhkEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExNDEzMzg0N1oXDTI0MTExMzEzMzg0N1owgY0xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEgMB4GA1UECgwXT25lU3RyZWFtIE5ldHdvcmtzLCBMTEMxGzAZBgNVBAsMEk9uZVN0cmVhbSBOZXR3b3JrczEsMCoGA1UEAwwjU0hBS0VOIE9uZVN0cmVhbSBOZXR3b3JrcywgTExDIDYzMEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASa11qvWcQ9DavvtESM1uR3AZMEpp%2BkFrLAxVgyINKUMeqH%2BkMfm5Y6%2FHw4EjwOJPLgclhrrvzU935N8NQqRNNVo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MzBKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUHprFro9TA3WtLfg3l%2FJJk7X4Yg4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDbCa6WKlhNebet0XqmnSFzirV%2BDiEdZ%2F%2Bbljd95HMr9gIhAKgq25LZQ40FlIzXO814pezMHXMq5MuliMvei0Idf89P)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 630J', but common name is 'SHAKEN OneStream Networks, LLC 630J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 24 Nov 23 11:17 UTC