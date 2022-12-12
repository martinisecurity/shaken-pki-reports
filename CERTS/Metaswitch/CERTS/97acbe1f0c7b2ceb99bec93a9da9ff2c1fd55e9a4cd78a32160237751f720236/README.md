# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Utility SHAKEN Cert 9262

Tested At: 12 Dec 22 23:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 565 day(s)\
Subject: CN=Utility SHAKEN Cert 9262, O=Utility Telecom Group\\, LLC, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/9da57471f2181f41af029f868fb29eb79e8af233

[View certificate details](https://understandingwebpki.com/?cert=MIICXTCCAgKgAwIBAgIQXRFPKpcKC%2FNDGSqa8LceTDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYzMDE2NTkzMVoXDTI0MDYyOTE2NTkzMVowVTELMAkGA1UEBhMCVVMxIzAhBgNVBAoMGlV0aWxpdHkgVGVsZWNvbSBHcm91cCwgTExDMSEwHwYDVQQDDBhVdGlsaXR5IFNIQUtFTiBDZXJ0IDkyNjIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATAYBp1RI39IZcnIrXQVxW3cQeI9EODlM39CgxGfwg%2Be6AjG7kmXUM82wG2qkgHOSU75qFlqk0T60b%2B6uaEJFc6o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5MjYyMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFOSJH0fVih1Cm%2BXNXQfEHfYyCVslMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQDz1cSgfLpK0BG%2Favq1aGFFM2HxQ6xRn8aBMlliFHZ81QIhAN%2BdzbLMPOHIFOrAZS1LoVh2qxO2YfZdcQhybacI34fU)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Dec 22 23:45 UTC