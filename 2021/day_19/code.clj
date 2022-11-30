(ns day-19.code
  (:require
    [clojure.math.combinatorics :refer
     [permutations combinations]]
    [clojure.set :refer [union intersection]]
    [clojure.string :as str]))


(def test-data
  (str/split-lines (slurp "day_19/test.txt")))

(def test-data-small
  (str/split-lines (slurp "day_19/test_small.txt")))

(defn build [lines]
  (:coll
    (reduce
      (fn [{:keys [idx] :as coll} n]
        (if-let [idx (last (re-find #"--- scanner (\d+).*" n))]
          (-> (assoc coll :idx (read-string idx))
              (assoc-in [:coll (read-string idx)] []))
          (update-in coll [:coll idx] conj (read-string (str "[" n "]")))))
      {:idx nil
       :coll {}}
      (remove #(= "" %) lines))))

(defn subtract [a b] (map - a b))
(defn add [a b] (map + a b))

(defn relset
  [coords]
  (reduce
    (fn [acc [a b]]
      (assoc acc (subtract a b) [a b])) {}
    (combinations coords 2)))

(defn perms
  "given 3 coordinates, split out into the 24 different variations"
  [coord]
  (mapcat
    (fn [[x y z]]
      [[(* -1 x) y        z       ]
       [x        (* -1 y) z       ]
       [x        y        (* -1 z)]
       [x        y        z       ]])
    (permutations coord)))

(defn all-perms
  [coords]
  (apply map vector (map perms coords)))

(def test-scanners
  (build test-data))

(defn find-similar [a b]
  (let [r1 (relset a)
        r2 (relset b)
        found (intersection (set (keys r1)) (set (keys r2)))]
    [(set (mapcat r1 found))
     (set (mapcat r2 found))]))

(defn overlapping-beacons
  [a b]
  (->> (all-perms b)
       (map (partial find-similar a))
       (map-indexed vector)
       (filter #(>= (count (first (second %))) 12))
       first
       ((fn [[idx [a b]]]
          (when idx
            {:perm-idx idx
             :a a
             :b b})))))

(defn beacons [scanners id1 id2]
  (overlapping-beacons (scanners id1) (scanners id2)))

(defn find-distance [a b]
  (first (intersection
    (set (map (partial add (first a)) b))
    (set (map (partial add (second a)) b)))))

(defn run [scanners]
  (loop
    [s (reduce
         (fn [acc [k v]]
           (assoc acc k {:d v}))
         {}
         scanners)
     curr 0
     l (disj (set (keys scanners)) 0)]
    (if (seq l)
      (let [found (->> (map #(overlapping-beacons (s curr) (s %)) l)
                       (remove nil?)
                       first)]
        found)
      s)))


(run test-scanners)

(let [s1 (test-scanners 0)
      s2 (test-scanners 1)
      {:keys [a b perm-idx]} (overlapping-beacons s1 s2)]
  (find-distance a b))

(let [[a b] (overlapping-beacons
              (map
                #(subtract [68 -1246 -43] %)
                (second (all-perms (test-scanners 1))))
              (test-scanners 3))]
  ;; (map (partial add [-20,-1133,1061]) b) (->>
    a
    ;; (map #(add % [-20,-1133,1061]))
    ;; (map #(subtract [68 -1246 -43] %))
    )
  ;; (= a (set (map (partial subtract [68 -1246 -43]) b)))
  ;; (intersection
  ;;   (set (map (partial add (first a)) b))
  ;;   (set (map (partial add (second a)) b)))
  )

(let [[a b] (overlapping-beacons (test-scanners 0)
                                 (test-scanners 1))]
  (= a (set (map (partial subtract [68 -1246 -43]) b)))
  ;; (intersection
  ;;   (set (map (partial add (first a)) b))
  ;;   (set (map (partial add (second a)) b)))

  )
