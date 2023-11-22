# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Number Access LLC SHAKEN 343J

Tested At: 22 Nov 23 03:32 UTC\
Initial Validity Period: 42 day(s)\
Remaining Validity Period: -220 day(s)\
Subject: L=Bettendorf, ST=Iowa, O=Number Access LLC, C=US, CN=Number Access LLC SHAKEN 343J\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://ss-public-certs.numberaccess.net/na-20230304154743-20230415154743.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICsDCCAjegAwIBAgIJAOY2%2FchuD0fhMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzAzMDQxNTQ3NDNaFw0yMzA0MTUxNTQ3NDNaMHUxJjAkBgNVBAMMHU51bWJlciBBY2Nlc3MgTExDIFNIQUtFTiAzNDNKMQswCQYDVQQGEwJVUzEaMBgGA1UECgwRTnVtYmVyIEFjY2VzcyBMTEMxDTALBgNVBAgMBElvd2ExEzARBgNVBAcMCkJldHRlbmRvcmYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATpL7xgZNdK%2FlaGTSWwLyiuhVCxCIZCM4T%2BuLIDuQae%2BuP2BOu6qu%2FI3arWpS1byYPLNQSruM4%2Bd8TIPOkygTDlo4GVMIGSMBYGCCsGAQUFBwEaBAowCKAGFgQzNDNKMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBRKOwGqCRp9XLRXQ37%2B3z7QlwbyjTAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDZwAwZAIwarldLQQuPj8aZCBXz9imAyt9zZYF2DlXezE1N6eEiqsLGtjYaV05yBkWp3rp7Zd0AjA0qpcLFA9bh7aPyeuVKdTeV3JvHnTqUxa0CfGITdMqTj1XRiaZhXuX9VD02iGG4iA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is marked as critical |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Nov 23 03:38 UTC