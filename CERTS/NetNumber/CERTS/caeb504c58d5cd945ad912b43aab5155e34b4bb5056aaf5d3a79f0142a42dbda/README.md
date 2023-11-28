# STIR/SHAKEN CA Ecosystem Compliance

## Certificate HD CARRIER LLC

Tested At: 28 Nov 23 20:12 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -88 day(s)\
Subject: L=Long Beach, ST=California, O=HD CARRIER LLC, C=US, CN=HD CARRIER LLC\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/0e78d59a-5421-4e8b-8d2f-37303891db4b

[View certificate details](https://understandingwebpki.com/?cert=MIICpTCCAiugAwIBAgIJAO%2B2TrNOAhUTMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzA4MDIwMDAwMDBaFw0yMzA4MzEyMzU5NTlaMGkxFzAVBgNVBAMMDkhEIENBUlJJRVIgTExDMQswCQYDVQQGEwJVUzEXMBUGA1UECgwOSEQgQ0FSUklFUiBMTEMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAcMCkxvbmcgQmVhY2gwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2BjoiWl6irYziXFAhlLCYvdJM89omNmggSAFtV%2F6DLW2LIqknZPy7jK3ki%2Fj1x9MSrlQ71x66y2t16Da5xroQto4GVMIGSMBYGCCsGAQUFBwEaBAowCKAGFgQzMjFKMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBRO4uQ%2Ffm72A4oT%2Fr6226RSEv8OhTAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIxAOz6aqMbNRH%2F%2Fc9xlW4%2BAvQNvZLQ8G0ZANyG3JYcPzrX4v5xeleVerwLb3oLcB%2BO%2BAIwEVeGc9Y3cwU2pv%2BAwf1BVGh26we4fhl6PTAKtHcEaCF2%2BxW2qsxFHdrtthwj0Aqw)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 321J', but common name is 'HD CARRIER LLC' |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'HD CARRIER LLC' does not contain 'SHAKEN' |


Generated: 28 Nov 23 20:21 UTC