#!/usr/bin/perl


use strict;

my($linein);


while ($linein = <STDIN>) {
    chomp($linein);

    $linein =~ s/^\s+"/ "/;
    $linein =~ s/^\s+}/ }/;
    $linein =~ s/^\s+]/ ]/;

    $linein =~ s/^}$/ }/;

    print(STDOUT $linein);

    if ($linein =~ /[[]$/)    { print("\n"); }
    if ($linein =~ /[}\]],$/) { print("\n"); }
}
print("\n");
