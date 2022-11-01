# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 747 certificates were included in the corpus being tested
- 108 repositories in the corpus were skipped because they are duplicates
- 305 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 332 certificates being tested against the remaining rules
- 93.98% of certificates contain one or more Error level issue
- 73.80% of certificates contain one or more Warning level issue
- 2.11% of certificates contain one or more Notice level issue
- 16.27% of certificates are too old to be assessed against currently enforced expectations
- 296 days is the average remaining validity for the certificates in the corpus
- 296 days is the average initial validity for the certificates in the corpus
- 164 certificates expire in the next 30 days

### CA Certificates

- 33 certificates were included in the corpus being tested
- 0 repositories in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 29 certificates being tested against the remaining rules
- 41.38% of certificates contain one or more Error level issue
- 86.21% of certificates contain one or more Warning level issue
- 62.07% of certificates contain one or more Notice level issue
- 75.86% of certificates are too old to be assessed against currently enforced expectations
- 5741 days is the average remaining validity for the certificates in the corpus
- 5630 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository

- 53.92% of certificate repositories contain one or more Error level issue
- 79.82% of certificates repositories contain one or more Warning level issue
- 0.00% of certificates repositories contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

## Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 43 (12.95%) | 43 (12.95%) | 43 (12.95%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 1 (0.30%) | 1 (0.30%) | 1 (0.30%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 7 (2.11%) | 0 (0.00%) | 0 (0.00%) | 4 (1.20%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 21 (6.33%) | 21 (6.33%) | 6 (1.81%) | 0 (0.00%) | 20 (6.02%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 3 (0.90%) | 3 (0.90%) | 1 (0.30%) | 3 (0.90%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 95 (28.61%) | 82 (24.70%) | 32 (9.64%) | 0 (0.00%) | 32 (9.64%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 5 (1.51%) | 5 (1.51%) | 5 (1.51%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 3 (0.90%) | 3 (0.90%) | 3 (0.90%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 91 (27.41%) | 91 (27.41%) | 91 (27.41%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.30%) | 1 (0.30%) | 1 (0.30%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 62 (18.67%) | 62 (18.67%) | 62 (18.67%) | 0 (0.00%) | 2 (0.60%) |
| **Total** | 332 (100.00%) | 312 (93.98%) | 245 (73.80%) | 7 (2.11%) | 54 (16.27%) |

## CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) | 1 (3.45%) | 2 (6.90%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (6.90%) | 1 (3.45%) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (10.34%) | 0 (0.00%) | 3 (10.34%) | 1 (3.45%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 2 (6.90%) | 2 (6.90%) | 0 (0.00%) | 0 (0.00%) | 2 (6.90%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (10.34%) | 2 (6.90%) | 3 (10.34%) | 0 (0.00%) | 3 (10.34%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (20.69%) | 3 (10.34%) | 6 (20.69%) | 6 (20.69%) | 4 (13.79%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (6.90%) | 1 (3.45%) | 2 (6.90%) | 2 (6.90%) | 1 (3.45%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) | 2 (6.90%) | 2 (6.90%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (6.90%) | 1 (3.45%) | 1 (3.45%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 0 (0.00%) | 2 (6.90%) | 2 (6.90%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 3 (10.34%) | 3 (10.34%) | 3 (10.34%) | 3 (10.34%) | 3 (10.34%) |
| **Total** | 29 (100.00%) | 12 (41.38%) | 25 (86.21%) | 18 (62.07%) | 22 (75.86%) |


Generated: 01/11/2022 at 07:33:04