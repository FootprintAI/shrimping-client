from time import sleep
import subprocess

def main():
    profilefile = 'testdata/profiles'
    postfile = 'testdata/posts'

    with open(profilefile, 'r') as f1:
        all_profiles = f1.read().splitlines()

    with open(postfile, 'r') as f1:
        all_shortcodes = f1.read().splitlines()

#    while True:
    for batchprofile in batch(all_profiles, 50):
        batchstr = ' '.join(batchprofile)
        cmd = "./instagram profile {} --svr_address 127.0.0.1:50050 --apikey 1c63f74e-685b-4f2c-ba2f-6adb276cb570".format(batchstr)
        exec(cmd)
        #sleep(600)

#    for batchprofile in batch(all_shortcodes, 50):
#        batchstr = ' '.join(batchprofile)
#        cmd = "./instagram profile {} --svr_address 127.0.0.1:50050 --apikey 1c63f74e-685b-4f2c-ba2f-6adb276cb570".format(batchstr)
#        exec(cmd)
#        sleep(3600)
#
#
#    for profile in all_profiles:
#        cmd = "./instagram posts {} --svr_address 127.0.0.1:50050 --apikey 1c63f74e-685b-4f2c-ba2f-6adb276cb570".format(profile)
#        exec(cmd)


def batch(iterable, n=1):
    l = len(iterable)
    for ndx in range(0, l, n):
        yield iterable[ndx:min(ndx + n, l)]

def exec(cmd):
    import os
    os.system(cmd)


main()
