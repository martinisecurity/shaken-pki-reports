# STIR/SHAKEN CA Ecosystem Compliance

## Certificate South Central Rural Telecommunications Coop SHAKEN Cert 0418

Tested At: 30 Nov 22 17:35 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 472 day(s)\
Subject: CN=South Central Rural Telecommunications Coop SHAKEN Cert 0418, O=South Central Rural Telecommunications Cooperative, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/99b4dd6734c3c1840bfdc0a2014214200e402920

[View certificate details](https://understandingwebpki.com/?cert=MIICmTCCAj%2BgAwIBAgIQEO%2BPCQIJ5L7otpfpK5jK7jAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxNzE3MTkxMVoXDTI0MDMxNjE3MTkxMVowgZExCzAJBgNVBAYTAlVTMTswOQYDVQQKDDJTb3V0aCBDZW50cmFsIFJ1cmFsIFRlbGVjb21tdW5pY2F0aW9ucyBDb29wZXJhdGl2ZTFFMEMGA1UEAww8U291dGggQ2VudHJhbCBSdXJhbCBUZWxlY29tbXVuaWNhdGlvbnMgQ29vcCBTSEFLRU4gQ2VydCAwNDE4MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAER2hU9qavCEnkQFWke5FaG%2FOVVgkGW2O6gaslY66VHeYObO3BwWsdknnGo5RXNF2Yp7LB9dhphhrRFkerisSbe6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMDQxODBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSlciGe1ch5Bc2QYGcUbjTdp4tSLjAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiABSWn6RmK4eYrnsAR9HVhh0gwGu5oeYOIjFxJTqD2hQgIhAKWqV5iHtE69jmA5kB7fn4dnR%2B9NTvgRu3vCKgJ2nOlq)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 30 Nov 22 17:39 UTC