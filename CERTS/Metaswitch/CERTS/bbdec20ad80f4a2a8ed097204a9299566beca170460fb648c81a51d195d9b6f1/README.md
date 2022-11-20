# STIR/SHAKEN CA Ecosystem Compliance

## Certificate RCN SHAKEN Cert 7615

Tested At: 20 Nov 22 22:56 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 482 day(s)\
Subject: CN=RCN SHAKEN Cert 7615, O=RCN, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/11d9e0e06eeaaf56fb9be54eec6aae5a8b3461fe

[View certificate details](https://understandingwebpki.com/?cert=MIICQTCCAeegAwIBAgIQCFfTK6Nnq8%2BDNtDYu%2BZYnzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxNzE3MDYwMVoXDTI0MDMxNjE3MDYwMVowOjELMAkGA1UEBhMCVVMxDDAKBgNVBAoMA1JDTjEdMBsGA1UEAwwUUkNOIFNIQUtFTiBDZXJ0IDc2MTUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATAt6ELjr2XfUM%2FjKVUUjVMaZcM45EZ%2FuxWR4jx0mAVh8Q6Uvfrxq8h8gnbLVtqhV%2FRMz0VNdsrqrG0WUpG4X3po4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NjE1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHd6jgfsvzke%2Bwfzg1VrzP25KaPeMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCICUcEjuJCsxEFC6zuTL7lSoVrTdYAyCDB%2Be8qnO4bol1AiEAtX1wCQcGbJmq4L9Hnu2WmV5xO1m4ccoKQHdKu52oYh0%3D)

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


Generated: 20 Nov 22 22:57 UTC