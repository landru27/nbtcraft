#!/usr/bin/perl


use strict;

my($linein);


while ($linein = <STDIN>) {
    chop($linein);
    chop($linein);

    $linein =~ s/^\s+"/ "/;
    $linein =~ s/^\s+}/ }/;
    $linein =~ s/^\s+]/ ]/;

    $linein =~ s/^}$/ }/;

    print(STDOUT $linein);

    if ($linein =~ /[[]$/)    { print("\r\n"); }
    if ($linein =~ /[}\]],$/) { print("\r\n"); }
}
print("\r\n");
