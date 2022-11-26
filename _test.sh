#!/bin/bash

# List the files that existed before the test
find . -type f > files_to_keep.txt

# Change into directory to run tests
cd sources

# Run the tests
go test -v -cover -coverprofile=coverage > ../TestResults/test_results.txt
go tool cover -html=coverage -o ../TestResults/coverage.html
mv coverage ../TestResults/coverage.out
cat ../TestResults/test_results.txt

# Generate badges
tests_failed=$(grep -E -i "FAIL:" ../TestResults/test_results.txt)
string_code_coverage=$(grep -oP "(?<=coverage: )\S+(?=\% of statements)" ../TestResults/test_results.txt)
int_code_coverage=$(printf '%.0f\n' $string_code_coverage)

# Test results badge
if $test_failed; then
    cp ../TestResults/Badges/Test_Results/Passed/green.svg ../TestResults/PassFailBadge.svg
else
    cp ../TestResults/Badges/Test_Results/Failed/red.svg ../TestResults/PassFailBadge.svg
fi

# Code Coverage badge
if [[ $int_code_coverage -ge 90 ]] ; then
    cp ../TestResults/Badges/Code_Coverage/green/"${int_code_coverage}_Percent.svg"  ../TestResults/CodeCoverageBadge.svg
elif [[ $int_code_coverage -ge 75 ]] ; then
    cp ../TestResults/Badges/Code_Coverage/yellow/"${int_code_coverage}_Percent.svg" ../TestResults/CodeCoverageBadge.svg
else
    cp ../TestResults/Badges/Code_Coverage/red/"${int_code_coverage}_Percent.svg"    ../TestResults/CodeCoverageBadge.svg
fi

# Remove unnecessary files
cd ..
find . -type f | grep -v -x -f files_to_keep.txt | xargs -d "\n" -P 0 rm -f

rm files_to_keep.txt