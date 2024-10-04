# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Encore Communications SHAKEN Cert 956H

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 301 day(s)\
Subject: CN=Encore Communications SHAKEN Cert 956H, O=Encore Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/7dd749ccd2d68458243cb752f8f94dc8d928d627

[View certificate details](https://x509.io/?cert=MIICZDCCAgugAwIBAgIQIK4LjinBRMxFLWYhmrJYljAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDgwMjEwNDMwM1oXDTI1MDgwMTEwNDMwM1owXjELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFUVuY29yZSBDb21tdW5pY2F0aW9uczEvMC0GA1UEAwwmRW5jb3JlIENvbW11bmljYXRpb25zIFNIQUtFTiBDZXJ0IDk1NkgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARI74cjIb%2BrPqf6GBpi8NPczFTMUT0hN9uOQjuWf85xGNoIdRyZYcoIOHqWFv1Io0N%2FeR2H%2BmlzX4WN4jQ57FzHo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5NTZIMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFF7kzE3d%2Bb3f1pZmSf9uIGaklFHDMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIAkxulNHgPzT%2Bo1qVduGj69gerGZHLl7aH7L9fs7f533AiARZFnht0V0pR8k119FfEoA9p5pTkrbpBfLvGIVvypXaQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 956H', but common name is 'Encore Communications SHAKEN Cert 956H' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC