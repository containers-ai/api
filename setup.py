#import os
#import sys

from setuptools import setup, find_packages

# Ensure we're in the proper directory whether or not we're being used by pip.
#os.chdir(os.path.dirname(os.path.abspath(__file__)))

version='0.1'

with open('README.md', 'r') as f:
    readme = f.read()


with open('LICENSE', 'r') as f:
    license = f.read()

#PACKAGE_DIRECTORIES = {
#    '': './aiservice/',
#}

INSTALL_REQUIRES = (
    'protobuf>=3.6.0',
    'grpcio>=1.6.0',
    'grpcio-tools>=1.6.0',
    'googleapis-common-protos>=1.5.3',
)

setup(
    name='alameda-api',
    version=version,
    description='Alameda API interfaces',
    long_description=readme,
    long_description_content_type="text/markdown",
    author='ProphetStor Inc.',
    author_email='support@prophetstor.com',
    urls='https://github.com/containers-ai/api',
    license=license,
    packages=['alameda_ai', 'alameda_operator'],
    package_dir={
        'ai_service.v1alpha1': 'aiservice/v1alpha1',
        'operator.v1alpha1': 'operator/v1alpha1'
    },
    install_requires=INSTALL_REQUIRES,
    zip_safe=False
)

