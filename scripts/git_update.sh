#!/bin/bash

VERSION=""

# get parameters
while getopts v: flag
do
    case "${flag}" in
        v) VERSION=${OPTARG};;
    esac
done

# get highest tag number, and add v0.1.0 if doesn't exist
git fetch --prune --unshallow 2>/dev/null
CURRENT_VERSION=`git describe --abbrev=0 --tags 2>/dev/null`

if [[ $CURRENT_VERSION == '' ]]; then
    CURRENT_VERSION='v0.1.0'
fi
echo "Current version: $CURRENT_VERSION"

# replace . with space, and split into array
CURRENT_VERSION_PARTS=(${CURRENT_VERSION//./ })

# get major, minor, patch
MAJOR=${CURRENT_VERSION_PARTS[0]}
MINOR=${CURRENT_VERSION_PARTS[1]}
PATCH=${CURRENT_VERSION_PARTS[2]}

if [[ $VERSION == 'major']]
then 
    MAJOR=v$((MAJOR+1))
elif [[ $VERSION == 'minor']]
then 
    MINOR=$((MINOR+1))
elif [[ $VERSION == 'patch']]
then 
    PATCH=$((PATCH+1))
else
    echo "Invalid version type"
    exit 1
fi

# create new tag
NEW_TAG="$MAJOR.$MINOR.$PATCH"
echo "($VERSION) updating $CURRENT_VERSION to $NEW_TAG"

# get current hash and see if it already has a tag
GIT_COMMIT=`git rev-parse HEAD`
NEEDS_TAG=`git describe --contains $GIT_COMMIT 2>/dev/null`

if [-z "$NEEDS_TAG"]; then
    echo "Tagged with $NEW_TAG"
    git tag $NEW_TAG
    git push --tags
    git push
else
    echo "Already a tag on this commit"
fi

echo ::set-output name=git-tag::$NEW_TAG