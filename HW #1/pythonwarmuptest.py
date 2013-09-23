#   Problem 10
from pythonwarmup import (change, strip_vowels, scramble, powers_of_two, powers, interleave, stutter)
import unittest

def anagram(s, t):
    return sorted(s) == sorted(t)

class TestGettingStartedFunctions(unittest.TestCase):

    def test_change(self):
        self.assertEqual((3, 2, 0, 2), change(97))
        self.assertEqual((1, 0, 1, 0), change(30))
        self.assertEqual((0, 0, 1, 3), change(8))
        self.assertEqual((5, 1, 0, 2), change(137))
        self.assertEqual((4, 0, 0, 0), change(100))
        self.assertEqual((0, 0, 0, 0), change(0))
        self.assertEqual((160103, 2, 0, 2), change(4002597))
        self.assertEqual((0, 0, 0, 4), change(4))
        with self.assertRaises(ValueError):
            change(-20)
        

    def test_strip_vowels(self):
        self.assertEqual('Hll, wrld', strip_vowels("Hello, world"))
        self.assertEqual('', strip_vowels("AEIOU"))
        self.assertEqual('', strip_vowels("aeiou"))
        self.assertEqual('Zn Qncy Knsl', strip_vowels("Zane Quincy Kansil"))
        self.assertEqual('Tm Mrvl Rddl', strip_vowels("Tom Marvolo Riddle"))
        self.assertEqual('BbCcDdFfGgHhJjKkLlMmNnPpQqRrSsTtVvWwXxYyZz', strip_vowels("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"))
        self.assertEqual('Blln lk  gt prh Wnfry mny~', strip_vowels("Ballin like I got Oprah Winfrey money~"))
        self.assertEqual('Swt hm lbm', strip_vowels("Sweet home Alabama"))
        self.assertEqual('nc pn  tm thr ws  bg hg gnt Mn p thr.', strip_vowels("Once upon a time there was a big huge giant Moon up there."))
        self.assertEqual('Ty Stry 2', strip_vowels("Toy Story 2"))
        self.assertEqual('Hrvst Mn, th rgn f fllng brcclsrs', strip_vowels("Harvest Moon, the reign of falling broccolisaurs"))
   
   
    def test_scramble(self):
        data = ["a", "rat", "BSOD", "BDFL", "Python testing", " a b c ", "Hannah Montana", "142452729"]
        for s in data:
            self.assertTrue(anagram(s, scramble(s)))


    def test_powers_of_two(self):
        self.assertEqual(tuple(powers_of_two(1)),(1,))
        self.assertEqual(tuple(powers_of_two(7)),(1,2,4))
        self.assertEqual(tuple(powers_of_two(8)),(1,2,4,8))
        self.assertEqual(tuple(powers_of_two(15)),(1,2,4,8))
        self.assertEqual(tuple(powers_of_two(16)),(1,2,4,8,16))
        self.assertEqual(tuple(powers_of_two(17)),(1,2,4,8,16))
        self.assertEqual(tuple(powers_of_two(31)),(1,2,4,8,16))
        self.assertEqual(tuple(powers_of_two(32)),(1,2,4,8,16,32))
        self.assertEqual(tuple(powers_of_two(33)),(1,2,4,8,16,32))
        self.assertEqual(tuple(powers_of_two(70)),(1,2,4,8,16,32,64))
        self.assertEqual(tuple(powers_of_two(130)),(1,2,4,8,16,32,64,128))
        self.assertEqual(tuple(powers_of_two(255)),(1,2,4,8,16,32,64,128))
        self.assertEqual(tuple(powers_of_two(256)),(1,2,4,8,16,32,64,128,256))
        self.assertEqual(tuple(powers_of_two(65536)),(1,2,4,8,16,32,64,128,256,512,1024,2048,4096,8192,16384,32768,65536))
        self.assertEqual(tuple(powers_of_two(-22)),())
        self.assertEqual(tuple(powers_of_two(0)),())


    def test_powers(self):
        self.assertEqual(tuple(powers(3,30)), (1,3,9,27))
        self.assertEqual(tuple(powers(3,27)), (1,3,9,27))
        self.assertEqual(tuple(powers(3,26)), (1,3,9))
        self.assertEqual(tuple(powers(4,64)), (1,4,16,64))
        self.assertEqual(tuple(powers(10,999)), (1,10,100))
        self.assertEqual(tuple(powers(10,1000)), (1,10,100,1000))
        self.assertEqual(tuple(powers(5,125)), (1,5,25,125))
        self.assertEqual(tuple(powers(5,126)), (1,5,25,125))
        self.assertEqual(tuple(powers(2,130)), (1,2,4,8,16,32,64,128))
        self.assertEqual(tuple(powers(3,1)), (1,))
        self.assertEqual(tuple(powers(3,3)), (1,3))
        self.assertEqual(tuple(powers(3,10)), (1,3,9))


    def test_interleave(self):
        self.assertEqual(interleave([1,2],[3,4]), [1,3,2,4])
        self.assertEqual(interleave([1,2],[3,4,5]), [1,3,2,4,5])
        self.assertEqual(interleave([1,2],[3,4,5,6]), [1,3,2,4,5,6])
        self.assertEqual(interleave(['a','b','c'],[3,4]), ['a',3,'b',4,'c'])
        self.assertEqual(interleave([1],[3]), [1,3])
        self.assertEqual(interleave([1,2],[]), [1,2])
        self.assertEqual(interleave(['Aqua','Start up'],[1,2]), ['Aqua',1,'Start up',2])
        self.assertEqual(interleave([1,2],['Aqua','Start up']), [1,'Aqua',2,'Start up'])
        self.assertEqual(interleave([1],[2,3,4,5,6,7,8,9]), [1,2,3,4,5,6,7,8,9])
        self.assertEqual(interleave([],[]), [])


    def test_stutter(self):
        self.assertEqual( stutter([]), [])
        self.assertEqual( stutter(['a','b','c']), ['a','a','b','b','c','c'])
        self.assertEqual( stutter(['a','a','a']), ['a','a','a','a','a','a'])
        self.assertEqual( stutter([1,2,3]), [1,1,2,2,3,3])
        self.assertEqual( stutter([[1],[2,3]]), [[1],[1],[2,3],[2,3]])
        self.assertEqual( stutter([[1,2,3]]), [[1,2,3],[1,2,3]])
        self.assertEqual( stutter([[1,2,[3]]]), [[1,2,[3]],[1,2,[3]]])
        self.assertEqual( stutter(['Repeat that']), ['Repeat that', 'Repeat that'])


if __name__ == '__main__':
    suite = unittest.TestLoader().loadTestsFromTestCase(TestGettingStartedFunctions)
    unittest.TextTestRunner(verbosity=2).run(suite)
