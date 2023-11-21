# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Utility SHAKEN Cert 9262

Tested At: 21 Nov 23 17:37 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 221 day(s)\
Subject: CN=Utility SHAKEN Cert 9262, O=Utility Telecom Group\\, LLC, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/40f8209b7db250928de2793d9848e67759fbd400

[View certificate details](https://understandingwebpki.com/?cert=MIICXTCCAgKgAwIBAgIQXRFPKpcKC%2FNDGSqa8LceTDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYzMDE2NTkzMVoXDTI0MDYyOTE2NTkzMVowVTELMAkGA1UEBhMCVVMxIzAhBgNVBAoMGlV0aWxpdHkgVGVsZWNvbSBHcm91cCwgTExDMSEwHwYDVQQDDBhVdGlsaXR5IFNIQUtFTiBDZXJ0IDkyNjIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATAYBp1RI39IZcnIrXQVxW3cQeI9EODlM39CgxGfwg%2Be6AjG7kmXUM82wG2qkgHOSU75qFlqk0T60b%2B6uaEJFc6o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5MjYyMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFOSJH0fVih1Cm%2BXNXQfEHfYyCVslMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQDz1cSgfLpK0BG%2Favq1aGFFM2HxQ6xRn8aBMlliFHZ81QIhAN%2BdzbLMPOHIFOrAZS1LoVh2qxO2YfZdcQhybacI34fU)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:53 UTC