(ns bob)

(defn- yell? [s]
  (and (re-find #"[a-zA-Z]+" s)
       (= (.toUpperCase s) s)))

(defn- question? [s]
  (.endsWith s "?"))

(defn- sentence? [s]
  (re-find #"\S+" s))

(defn response-for [s]
  (cond
    (yell? s) "Whoa, chill out!"
    (question? s) "Sure."
    (sentence? s) "Whatever."
    :else "Fine. Be that way!"))