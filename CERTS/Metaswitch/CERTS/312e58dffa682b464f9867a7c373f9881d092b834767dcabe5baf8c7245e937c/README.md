# STIR/SHAKEN CA Ecosystem Compliance

## Certificate GCI SHAKEN Cert 7785

Tested At: 23 Nov 22 18:08 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 479 day(s)\
Subject: CN=GCI SHAKEN Cert 7785, O=GCI, C=us\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/cd1856717574765eb6b4bddb7a5bc8814e1e2103

[View certificate details](https://understandingwebpki.com/?cert=MIICQjCCAeegAwIBAgIQWVIfIucXyRwdk2r1QxfxijAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxNzE3MTI0MFoXDTI0MDMxNjE3MTI0MFowOjELMAkGA1UEBhMCdXMxDDAKBgNVBAoMA0dDSTEdMBsGA1UEAwwUR0NJIFNIQUtFTiBDZXJ0IDc3ODUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARaB2tkqnNVmOGuizyDsUBBLFhwWwoHf7WBK%2B539zfhV1dT5OQrbZJOgmiMy1cW7q%2BOAd1uuyOeag%2B2jcvLXmquo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3Nzg1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFNBTKUsbL5mmfkdrpXdVRnf1NlU8MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQD1plr6Dl%2FKnrzzFBGkxc17Py0p0BToNjrqn8ReVd9RWQIhAK8BHwvgTKCJ8QAi2yCwJN%2Bc%2BWK4Jd2oKPkmxZuKlwHL)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 23 Nov 22 18:09 UTC