#!/usr/bin/perl

# A Brighter Place: 

# This program will generate a positive message and a random quote from
# the web to help brighten up your day.

use strict;
use warnings;

use LWP::Simple;
use JSON::PP;

# Get a random positive message from the web
my $url = 'http://quotesondesign.com/wp-json/posts?filter[orderby]=rand&filter[posts_per_page]=1';
my $json_string = get($url);
my $data = decode_json $json_string;
my $quote = $data->[0]->{'content'};

# Clean up HTML content from the quote content
$quote =~ s/<[^>]*>//g;

# Generate Positive Message
my @positive_messages = (
    "You are capable of wonderful things!",
    "Today is a great day for something awesome!",
    "Every day is an opportunity for something new!",
    "Be the best version of yourself!",
    "Be the best you can be and always stay positive!"
);
my $random_index = int(rand(@positive_messages));
my $positive_message = $positive_messages[$random_index];

#Print Message
print "\n$positive_message\n$quote\n\n";