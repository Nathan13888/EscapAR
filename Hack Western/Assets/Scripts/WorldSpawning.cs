using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.XR.ARFoundation;

public class WorldSpawning : MonoBehaviour
{
    public bool spawned;
    public ARRaycastManager raycastManager;
    public GameObject world;
    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        if (!spawned){
            List<ARRaycastHit> hits = new List<ARRaycastHit>();
            raycastManager.Raycast(Vector3.down, hits, UnityEngine.XR.ARSubsystems.TrackableType.Planes);
            if (hits.Count > 0)
                {
                    GameObject.Instantiate(world, hits[0].pose.position, hits[0].pose.rotation);
                    spawned = true;
                }
        }
    }
}
