# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

445 certificates were included in the corpus being tested\
301 certificates in the corpus were skipped because they were expired\
2 certificates in the corpus were skipped because they are not currently trusted\
89.21% of certificates contain one or more Error level issue\
73.03% of certificates contain one or more Warning level issue\
1.80% of certificates contain one or more Notice level issue\
31.24% of certificates are too old to be assessed against currently enforced expectations\
429 days is the average remaining validity for the certificates in the corpus\
430 days is the average initial validity for the certificates in the corpus\
171 certificates expire in the next 30 days

### CA Certificates

29 certificates were included in the corpus being tested\
2 certificates in the corpus were skipped because they were expired\
2 certificates in the corpus were skipped because they are not currently trusted\
41.38% of certificates contain one or more Error level issue\
86.21% of certificates contain one or more Warning level issue\
62.07% of certificates contain one or more Notice level issue\
75.86% of certificates are too old to be assessed against currently enforced expectations\
5741 days is the average remaining validity for the certificates in the corpus\
5630 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

## Certificate Repository

64.27% of certificate repositories contain one or more Error level issue\
83.82% of certificates repositories contain one or more Warning level issue\
0.00% of certificates repositories contain one or more Notice level issue

## Details

## Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 43 (9.66%) | 43 (9.66%) | 43 (9.66%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 1 (0.22%) | 1 (0.22%) | 1 (0.22%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 7 (1.57%) | 0 (0.00%) | 0 (0.00%) | 4 (0.90%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 21 (4.72%) | 21 (4.72%) | 6 (1.35%) | 0 (0.00%) | 20 (4.49%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 4 (0.90%) | 4 (0.90%) | 1 (0.22%) | 4 (0.90%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 200 (44.94%) | 159 (35.73%) | 105 (23.60%) | 0 (0.00%) | 117 (26.29%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 5 (1.12%) | 5 (1.12%) | 5 (1.12%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 3 (0.67%) | 3 (0.67%) | 3 (0.67%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 94 (21.12%) | 94 (21.12%) | 94 (21.12%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.22%) | 1 (0.22%) | 1 (0.22%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 66 (14.83%) | 66 (14.83%) | 66 (14.83%) | 0 (0.00%) | 2 (0.45%) |
| **Total** | 445 (100.00%) | 397 (89.21%) | 325 (73.03%) | 8 (1.80%) | 139 (31.24%) |

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


Generated: 31/10/2022 at 19:21:49